package account

import (
	"testing"
)

func TestAccount_ValidateTo(t *testing.T) {
	type fields struct {
		Type Type
	}
	type args struct {
		Account Account
		WantErr bool
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
	}{
		{
			name:   "Should validate SalaryIncoming against other types",
			fields: fields{Type: SalaryIncoming},
			args: []args{
				{
					Account: Account{Type: SalaryIncoming},
					WantErr: true,
				},
				{
					Account: Account{Type: RevenueIncoming},
					WantErr: true,
				},
				{
					Account: Account{Type: InvestmentAccount},
					WantErr: true,
				},
				{
					Account: Account{Type: CheckingAccount},
					WantErr: false,
				},
				{
					Account: Account{Type: CreditCard},
					WantErr: true,
				},
				{
					Account: Account{Type: Loan},
					WantErr: true,
				},
				{
					Account: Account{Type: FixedAsset},
					WantErr: true,
				},
				{
					Account: Account{Type: Service},
					WantErr: true,
				},
				{
					Account: Account{Type: Consumable},
					WantErr: true,
				},
				{
					Account: Account{Type: Depreciation},
					WantErr: true,
				},
				{
					Account: Account{Type: Fee},
					WantErr: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Account{
				Type: tt.fields.Type,
			}

			for _, arg := range tt.args {
				if err := a.ValidateTo(arg.Account); (err != nil) != arg.WantErr {
					t.Errorf("ValidateTo() error = %v, wantErr %v", err, arg.WantErr)
				}
			}
		})
	}
}
