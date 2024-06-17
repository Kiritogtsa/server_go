package vendedor

import (
	"database/sql"
	"errors"
	"log"

	"server/src/models"
)

const Dns string = "root:1234@tcp(127.0.0.1:3306)/loja"

// Vendedordao representa o DAO para a entidade Vendedor
type Vendedordao struct {
	Conn *models.Conn
}

// NewVendedordao cria uma nova instância do DAO Vendedordao
func NewVendedordao() (*Vendedordao, error) {
	conn, err := models.NewConn(Dns)
	if err != nil {
		return nil, errors.New("erro ao criar a conexão com o banco de dados")
	}
	return &Vendedordao{Conn: conn}, nil
}

// Insert insere um novo vendedor no banco de dados
func (coon *Vendedordao) Insert(vendedor *Vendedor) (*Vendedor, error) {
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}
	query := "INSERT INTO vendedores (usuario_id, produtos) VALUES (?, ?)"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(vendedor.GetuserId(), vendedor.Getquant())
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Erro ao obter o último ID inserido:", err)
		return nil, err
	}

	vendedor.SetId(int(id))
	return vendedor, nil
}

// Update atualiza um vendedor existente no banco de dados
func (conn *Vendedordao) Update(vendedor *Vendedor) error {
	dao := conn.Conn.Getdb()
	if dao == nil {
		return nil
	}
	query := "UPDATE vendedores SET usuario_id = ?, produtos = ? WHERE id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vendedor.GetuserId(), vendedor.Getquant(), vendedor.GetId())
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return err
	}

	return nil
}

// Persist persiste o vendedor no banco de dados (insere ou atualiza dependendo da existência do ID)
func (dao *Vendedordao) Persist(vendedor *Vendedor) (*Vendedor, error) {
	if vendedor.GetId() == 0 {
		return dao.Insert(vendedor)
	}
	err := dao.Update(vendedor)
	if err != nil {
		return nil, err
	}
	return vendedor, nil
}

// FindById retorna um vendedor com base no ID
func (coon *Vendedordao) FindById(id int) (*Vendedor, error) {
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}
	query := "SELECT id, usuario_id, produtos FROM vendedores WHERE id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	vendedor := &Vendedor{}
	err = stmt.QueryRow(id).Scan(&vendedor.id, &vendedor.user, &vendedor.produtos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	return vendedor, nil
}
func (coon *Vendedordao) FindByUserid(id int) (*Vendedor, error) {
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}
	query := "SELECT id, usuario_id, produtos FROM vendedores WHERE usuario_id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	vendedor := &Vendedor{}
	err = stmt.QueryRow(id).Scan(&vendedor.id, &vendedor.user, &vendedor.produtos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	return vendedor, nil
}

// Delete remove um vendedor com base no ID
func (coon *Vendedordao) Delete(id int) error {
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil
	}
	query := "DELETE FROM vendedores WHERE id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return err
	}

	return nil
}
