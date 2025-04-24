package storage

import (
    "os"
    "testing"
	"CLIBookmark/internal/model"
    "github.com/stretchr/testify/assert"
)

func TestJSONStorage_AddAndGetAll(t *testing.T) {
    // 1. Создаём временный файл
    tmpFile, err := os.CreateTemp("", "test-bookmarks-*.json")
    assert.NoError(t, err)
    defer os.Remove(tmpFile.Name()) // Удаляем после теста

    // 2. Инициализируем хранилище
    storage := &JSONStorage{FilePath: tmpFile.Name()}

    // 3. Добавляем закладку
    err = storage.AddBookmark(model.Bookmark{
        ID:   1,
        Name: "Test",
        URL:  "https://example.com",
    })
    assert.NoError(t, err)

    // 4. Проверяем, что закладка сохранилась
    bookmarks, err := storage.GetAllBookmarks()
    assert.NoError(t, err)
    assert.Equal(t, 1, len(bookmarks))
    assert.Equal(t, "Test", bookmarks[0].Name)
}

func TestJSONStorage_Delete(t *testing.T) {
    // Аналогично, но проверяем удаление
}
