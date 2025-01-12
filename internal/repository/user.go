package repository

import (
	"OSS-Matching-ServerSide/internal/model"

	"gorm.io/gorm"
)

// インターフェース定義
type UserRepository interface {
	Create(db *gorm.DB, user *model.User) (*model.User, error)
	Get(db *gorm.DB, id string) (*model.User, error)
	Update(db *gorm.DB, user *model.User) error
	Delete(db *gorm.DB, id string) error
}

// 構造体と初期化
// 空の構造体userRepositoryを定義し、インターフェースを実装
type userRepository struct{}

// NewUserRepository()はUserRepositoryインターフェースを実装した新しいインスタンスを返す
// UserRepositoryインターフェースを実装した新しいインスタンスを作成するファクトリー関数です。
// 戻り値の型がUserRepository（インターフェース）になっている点が重要
// 実際には&userRepository{}（構造体のポインタ）を返すが、インターフェースとして扱える
// これにより、呼び出し側は具体的な実装（userRepository構造体）を知る必要がなく、インターフェースを通じて操作できる
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// (r *userRepository): メソッドレシーバー
// - *userRepository: このメソッドが属する構造体型へのポインタ。これによってこの関数はuserRepository構造体のメソッドとして定義される。
// 引数
// - db *gorm.DB: データベースコネクションのポインタ
// *はポインタを示し、DBインスタンスへの参照を渡す
// 詳細はhttps://zenn.dev/articles/3e2f3bdc378f19/editでまとめる
// - user *model.User: 作成するユーザーデータのポインタ
// model.Userは別パッケージで定義されたUser構造体
// ポインタを使用することでメモリ効率が良く、データの変更が可能
func (r *userRepository) Create(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Get(db *gorm.DB, id string) (*model.User, error) {
	var user model.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(db *gorm.DB, user *model.User) error {
	return db.Save(user).Error
}

func (r *userRepository) Delete(db *gorm.DB, id string) error {
	return db.Delete(&model.User{}, "id = ?", id).Error
}
