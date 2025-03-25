package main

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
)

// ファイルのデータを読み込む
func readFileContent(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ファイルにデータを書き込む
func writeFileContent(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}

// MarkdownファイルをHTMLに変換し、出力ファイルに書き込む
func convertMarkdownToHTML(inputPath, outputPath string) error {
	md, err := readFileContent(inputPath)
	if err != nil {
		return fmt.Errorf("ファイルの読み込みに失敗しました: %w", err)
	}

	// Markdown を HTML に変換
	html := markdown.ToHTML(md, nil, nil)

	if err := writeFileContent(outputPath, html); err != nil {
		return fmt.Errorf("ファイルの書き込みに失敗しました: %w", err)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) != 4 || args[1] != "markdown" {
		fmt.Println("使い方: go run main.go markdown <入力ファイル> <出力ファイル>")
		return
	}

	inputPath := args[2]
	outputPath := args[3]

	if err := convertMarkdownToHTML(inputPath, outputPath); err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println("変換が完了しました。")
}