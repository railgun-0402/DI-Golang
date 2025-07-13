package domain

import (
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestNewProduct_ValidInput(t *testing.T) {
	id := ulid.Make().String()
	tests := []struct {
		name        string
		productID   string
		nameInput   string
		descInput   string
		priceInput  int64
		stockInput  int
		wantErr     bool
		errContains string
	}{
		{
			name:       "valid input",
			productID:  id,
			nameInput:  "テスト商品",
			descInput:  "これはテスト商品です",
			priceInput: 1000,
			stockInput: 10,
			wantErr:    false,
		},
		{
			name:        "empty name",
			productID:   id,
			nameInput:   "",
			descInput:   "説明",
			priceInput:  1000,
			stockInput:  10,
			wantErr:     true,
			errContains: "商品名を入力してください",
		},
		{
			name:        "name too long",
			productID:   id,
			nameInput:   string(make([]rune, MaxNameLength+1)),
			descInput:   "説明",
			priceInput:  1000,
			stockInput:  10,
			wantErr:     true,
			errContains: "商品名は100文字以下",
		},
		{
			name:        "empty description",
			productID:   id,
			nameInput:   "商品",
			descInput:   "",
			priceInput:  1000,
			stockInput:  10,
			wantErr:     true,
			errContains: "商品の説明を入力してください",
		},
		{
			name:        "description too long",
			productID:   id,
			nameInput:   "商品",
			descInput:   string(make([]rune, MaxDescriptionLength+1)),
			priceInput:  1000,
			stockInput:  10,
			wantErr:     true,
			errContains: "商品説明は1000文字以下",
		},
		{
			name:        "invalid price",
			productID:   id,
			nameInput:   "商品",
			descInput:   "説明",
			priceInput:  0,
			stockInput:  10,
			wantErr:     true,
			errContains: "価格の値が不正です",
		},
		{
			name:        "invalid stock",
			productID:   id,
			nameInput:   "商品",
			descInput:   "説明",
			priceInput:  1000,
			stockInput:  -1,
			wantErr:     true,
			errContains: "在庫数の値が不正です",
		},
	}

	for _, tt := range tests {
		tt:= tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p, err := newProduct(
				tt.productID,
				tt.nameInput,
				tt.descInput,
				tt.priceInput,
				tt.stockInput,
			)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				if !contains(err.Error(), tt.errContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errContains, err.Error())
				}
				if p != nil {
					t.Errorf("expected product to be nil when error, got %+v", p)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if p == nil {
					t.Fatal("expected product but got nil")
				}
				if p.ID != tt.productID {
					t.Errorf("expected ID %q, got %q", tt.productID, p.ID)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) > 0 && string(s[:len(substr)]) == substr || len(s) > len(substr) && contains(s[1:], substr))
}

func TestNewProduct_InvalidName(t *testing.T) {
	id := ulid.Make().String()
	description := "This is test Product"
	price := int64(10000)
	stock := 10

	product, err := newProduct(id, "", description, price, stock)
	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}

	expected := "商品名を入力してください。"
	if err.Error() != expected {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}

	if product != nil {
		t.Errorf("expected product to be nil, got %+v", product)
	}
}

func TestNewProduct_InvalidNameMaxLength(t *testing.T) {

}
