package i18nx

import (
	"bufio"
	"os"
	"strings"
)

type I18nx struct {
	bundlePath string
	content    map[string]map[string]string
}

// New 创建 I18nx 实例
func New(bundlePath string) (*I18nx, error) {
	var (
		err     error
		entries []os.DirEntry
		m       map[string]string
	)

	if entries, err = os.ReadDir(bundlePath); err != nil {
		return nil, err
	}

	i := I18nx{
		bundlePath: bundlePath,
		content:    make(map[string]map[string]string),
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if m, err = resolveBundle(bundlePath + string(os.PathSeparator) + entry.Name()); err != nil {
			return nil, err
		}

		i.content[entry.Name()] = m
	}

	return &i, nil
}

// resolveBundle 解析 bundle 文件中的内容，存入 map 中
func resolveBundle(path string) (map[string]string, error) {
	var (
		err   error
		file  *os.File
		lines []string
	)

	if file, err = os.Open(path); err != nil {
		return nil, err
	}

	if lines, err = readLines(file); err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for _, line := range lines {
		if !strings.Contains(line, "=") {
			continue
		}
		arr := strings.Split(line, "=")
		m[arr[0]] = arr[1]
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return m, nil
}

// Translate 将 bundle 编码翻译成对应语言的内容
func (i *I18nx) Translate(i18nCode string, lang string) string {
	return i.content[lang][i18nCode]
}

// readLines 读取 file 中的所有行
func readLines(file *os.File) ([]string, error) {
	r := bufio.NewReader(file)

	var lines []string
	for {
		line, err := readLine(r)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return lines, nil
}

// readLine 读取 r 中的下一行
func readLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line     []byte
		ln       []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}
