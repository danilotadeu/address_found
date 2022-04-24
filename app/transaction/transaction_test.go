package transaction

import (
	"context"
	"errors"
	"testing"

	accountModel "github.com/danilotadeu/pismo/model/account"
	"github.com/danilotadeu/pismo/model/transaction"
)

func TestCreateTransaction(t *testing.T) {
	type args struct {
		ctx             context.Context
		transactionBody transaction.TransactionRequest
	}
	var ID int64
	ID = 1
	tests := []struct {
		name              string
		args              args
		want              *int64
		wantErr           bool
		GetAccount        func(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error)
		CreateTransaction func(ctx context.Context, accountID int64, operationTypeID int, amount float64) (*int64, error)
	}{
		{
			name: "Should return transaction invalid",
			args: args{
				ctx: context.Background(),
				transactionBody: transaction.TransactionRequest{
					AccountID:       1,
					OperationTypeID: 10,
					Amount:          1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should return get account not found",
			args: args{
				ctx: context.Background(),
				transactionBody: transaction.TransactionRequest{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          1,
				},
			},
			want:    nil,
			wantErr: true,
			GetAccount: func(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error) {
				return nil, accountModel.ErrorAccountNotFound
			},
		},
		{
			name: "Should return get account error",
			args: args{
				ctx: context.Background(),
				transactionBody: transaction.TransactionRequest{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          1,
				},
			},
			want:    nil,
			wantErr: true,
			GetAccount: func(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error) {
				return nil, errors.New("Internal server error")
			},
		},
		{
			name: "Should return error to create transaction",
			args: args{
				ctx: context.Background(),
				transactionBody: transaction.TransactionRequest{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          1,
				},
			},
			want:    nil,
			wantErr: true,
			GetAccount: func(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error) {
				return &accountModel.AccountResultQuery{
					ID:             1,
					DocumentNumber: "12345678910",
				}, nil
			},
			CreateTransaction: func(ctx context.Context, accountID int64, operationTypeID int, amount float64) (*int64, error) {
				return nil, errors.New("Internal server error")
			},
		},
		{
			name: "Should return create transaction",
			args: args{
				ctx: context.Background(),
				transactionBody: transaction.TransactionRequest{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          1,
				},
			},
			want:    &ID,
			wantErr: false,
			GetAccount: func(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error) {
				return &accountModel.AccountResultQuery{
					ID:             1,
					DocumentNumber: "12345678910",
				}, nil
			},
			CreateTransaction: func(ctx context.Context, accountID int64, operationTypeID int, amount float64) (*int64, error) {

				return &ID, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAccount = tt.GetAccount
			CreateTransaction = tt.CreateTransaction
			a := &appImpl{}
			got, err := a.CreateTransaction(tt.args.ctx, tt.args.transactionBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("appImpl.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("appImpl.CreateTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
