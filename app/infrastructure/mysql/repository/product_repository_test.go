package repository

import (
	"context"
	"log"
	"reflect"
	"regexp"
	"testing"

	"github.com/o-ga09/tutorial-go-fr/app/domain/product"
	"github.com/o-ga09/tutorial-go-fr/pkg/dbmock"
	"gorm.io/gorm"
)

func Test_productRepository_FindByID(t *testing.T) {
	// *gorm,DBのモックを生成
	db, mockdb, err := dbmock.GetNewDbMock()
	if err != nil {
		log.Fatal(err)
	}

	// クエリが実行された後の結果を設定
    rows := mockdb.NewRows([]string{"id","owner_id","name","description","price","stock"}).
                    AddRow("001","test","user1","test","1","1")
	mockdb.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `id` = ? ORDER BY `products`.`id` LIMIT 1")).WithArgs("001").WillReturnRows(rows)
	repository := NewProductRepository(db)

	type args struct {
		id string
	}
	type fields struct {
		id          string // 商品ID
		ownerID     string // 出品者ID
		name        string // 商品名
		description string // 商品の説明
		price       int64  // 商品金額
		stock       int    // 商品在庫
	}
	tests := []struct {
		name    string
		args args
		want    fields
		wantErr bool
	}{
		{name: "正常系",args: args{id: "001"},want: fields{id: "001",ownerID: "test",name: "user1",description: "test",price: 1,stock: 1},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := product.Reconstruct(tt.want.id,tt.want.ownerID,tt.want.name,tt.want.description,tt.want.price,tt.want.stock)
			got, err := repository.FindByID(context.Background(),tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("productRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productRepository_FindByIDs(t *testing.T) {
	// *gorm,DBのモックを生成
	db, mockdb, err := dbmock.GetNewDbMock()
	if err != nil {
		log.Fatal(err)
	}

	// クエリが実行された後の結果を設定
    rows := mockdb.NewRows([]string{"id","owner_id","name","description","price","stock"}).
                    AddRow("001","test","user1","test","1","1").
					AddRow("002","test","user2","test","1","1").
					AddRow("003","test","user3","test","1","1")
	mockdb.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE id IN (?,?,?)")).WithArgs("001","002","003").WillReturnRows(rows)
	repository := NewProductRepository(db)

	type args struct {
		ids []string
	}
	type fields struct {
		id          string // 商品ID
		ownerID     string // 出品者ID
		name        string // 商品名
		description string // 商品の説明
		price       int64  // 商品金額
		stock       int    // 商品在庫
	}
	tests := []struct {
		name    string
		args args
		want    []fields
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{[]string{"001","002","003"}},
			want: []fields{
				{id: "001",ownerID: "test",name: "user1",description: "test",price: 1,stock: 1},
				{id: "002",ownerID: "test",name: "user2",description: "test",price: 1,stock: 1},
				{id: "003",ownerID: "test",name: "user3",description: "test",price: 1,stock: 1},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := []product.Product{}
			for _, f := range tt.want {
				r, _ := product.Reconstruct(f.id,f.ownerID,f.name,f.description,f.price,f.stock)
				want = append(want, *r)
			}
			got, err := repository.FindByIDs(context.Background(),tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("productRepository.FindByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("productRepository.FindByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productRepository_Save(t *testing.T) {
	// Saveのテストをスキップする
	t.Skip()
}

func Test_productRepository_Update(t *testing.T) {
	// Updateのテストはスキップする
	t.Skip()
}

func Test_productRepository_Delete(t *testing.T) {
	// Deleteのテストはスキップする
	t.Skip()
}
func TestNewProductRepository(t *testing.T) {
	db, _, _ := dbmock.GetNewDbMock()
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want product.ProductRepository
	}{
		{name: "正常系",args: args{conn: db},want: NewProductRepository(db)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
