package product

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/o-ga09/tutorial-go-fr/pkg/strings"
)

func TestNewProduct(t *testing.T) {
	type args struct {
		ownerID     string
		name        string
		description string
		price       int64
		stock       int
	}

	cases := []struct {
		name    string
		args    args
		want    *Product
		wanrErr bool
	}{
		{name: "正常系", args: args{ownerID: "0000000001", name: "test1", description: "test1", price: 100, stock: 0}, want: &Product{ownerID: "0000000001", name: "test1", description: "test1", price: 100, stock: 0}, wanrErr: false},
		{name: "正常系", args: args{ownerID: "0000000002", name: "test2", description: "test2", price: 100, stock: 0}, want: &Product{ownerID: "0000000002", name: "test2", description: "test2", price: 100, stock: 0}, wanrErr: false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduct(tt.args.ownerID, tt.args.name, tt.args.description, tt.args.price, tt.args.stock)
			if (err != nil) != tt.wanrErr {
				t.Errorf("NewProduct() error = %v, wantErr %v ", err, tt.wanrErr)
				return
			}

			diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(Product{}), cmpopts.IgnoreFields(Product{}, "id"))
			if diff != "" {
				t.Errorf("NewProduct() = %v, want %v, error is %s, ", got, tt.want, err)
			}
		})
	}
}

func TestProduct_Consume(t *testing.T) {
	product := Product{
		id: strings.RemoveHyphen(uuid.New().String()),
		ownerID: strings.RemoveHyphen(uuid.New().String()),
		name: "test",
		description: "test",
		price: 100,
		stock: 100,
	}

	type fields struct {
		id          string
		ownerID     string
		name        string
		description string
		price       int64
		stock       int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "正常系",fields: fields(product),args: args{quantity: 10},wantErr: false},
		{name: "在庫が足りない",fields: fields(product),args: args{quantity: 101},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				id:          tt.fields.id,
				ownerID:     tt.fields.ownerID,
				name:        tt.fields.name,
				description: tt.fields.description,
				price:       tt.fields.price,
				stock:       tt.fields.stock,
			}
			if err := p.Consume(tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("Product.Consume() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
