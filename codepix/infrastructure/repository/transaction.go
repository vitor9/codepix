package repository

import (
	"fmt"
	"github.com/vitor9/codepix/tree/master/codepix/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}
// Register cria um novo registro
func (t* TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error

	if err != nil {
		return err
	}
	
	return nil
}

// Save, salva um registro, update. Se realizei alguma alteracao num model, vamos dar um save chamando esse repo
// Pasasndo por parametro esse model que alteramos
func (t* TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	// Preload para saber de onde veio a transacao, e qual o banco. Pq se nao ia retornar soh o account from id
	// Com isso ele vai trazer todas as informacoes da conta e do banco
	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction found")
	}
	return &transaction, nil
}