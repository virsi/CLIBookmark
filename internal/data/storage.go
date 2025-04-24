package storage

import (
	"os"
	"errors"
	"encoding/json"
	"CLIBookmark/internal/model"
)

type JSONStorage struct {
	FilePath string
}

func (s* JSONStorage) GetAllBookmarks() ([]model.Bookmark, error) {
	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, errors.New("no file in this path")
	}
	var bookmarks []model.Bookmark
	json.Unmarshal(data, &bookmarks)
	return bookmarks, nil
}

func (s *JSONStorage) AddBookmark(bookmark model.Bookmark) error {
    bookmarks, _ := s.GetAllBookmarks()
    bookmarks = append(bookmarks, bookmark)
    data, _ := json.MarshalIndent(bookmarks, "", "  ")
    return os.WriteFile(s.FilePath, data, 0644)
}

func (s *JSONStorage) DeleteBookmark(bookmark model.Bookmark) error {
    bookmarks, _ := s.GetAllBookmarks()
    var updatedBookmarks []model.Bookmark
    for i, bm := range bookmarks {
        if bm.Name == bookmark.Name {
            updatedBookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
            break
        }
    }

    if updatedBookmarks == nil {
        return errors.New("bookmark not found")
    }

    data, err := json.MarshalIndent(updatedBookmarks, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(s.FilePath, data, 0644)
}
