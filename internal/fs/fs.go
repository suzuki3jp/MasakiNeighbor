package fs

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/suzuki3jp/mn/internal/entity"
)

// ErrNotCSV は指定されたパスの拡張子が .csv でない場合に返されます。
var ErrNotCSV = errors.New("CSV: 指定されたパスが .csv ファイルではありません")

// InvalidRowError は CSV の特定行の解析で見つかったエラーを表します。
// Line はファイル上の 1 始まりの行番号（エラー発生位置）です。
type InvalidRowError struct {
	Line int
	Msg  string
}

func (e *InvalidRowError) Error() string {
	return fmt.Sprintf("CSV: %d 行目に不正な値: %s", e.Line, e.Msg)
}

// ReadPointsCsv はファイルパスの CSV を読み込み、各データ行を entity.Point に変換して返します。
// CSV はヘッダ行を持ち（スキップされます）、列の順序は x,y,series を想定します。
// 拡張子が .csv でない場合は ErrNotCSV を返します。
// データ行のいずれかに不正がある（パースエラー、series が期待値でない等）場合、
// 該当する行番号を含む *InvalidRowError を返します。
func ReadPointsCsv(path string) ([]entity.Point, error) {
	if filepath.Ext(path) != ".csv" {
		return nil, ErrNotCSV
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	// ParseRecords はファイル IO を伴わない純粋関数で、テスト可能
	return ParseRecords(records)
}

// --- 以降は純粋関数: ファイル IO を伴わず、テストしやすい ---

// ParseRecord は CSV の 1 行分（x,y,series）を解析して entity.Point を返す。
// 引数 lineNo はファイル上の 1-origin の行番号（エラー報告用）。
func ParseRecord(lineNo int, rec []string) (entity.Point, error) {
	if len(rec) < 3 {
		return entity.Point{}, &InvalidRowError{Line: lineNo, Msg: fmt.Sprintf("列が足りません: %d 列", len(rec))}
	}
	xstr := rec[0]
	ystr := rec[1]
	series := rec[2]

	x, err := strconv.ParseFloat(xstr, 64)
	if err != nil {
		return entity.Point{}, &InvalidRowError{Line: lineNo, Msg: fmt.Sprintf("x の数値変換に失敗: %v", err)}
	}
	y, err := strconv.ParseFloat(ystr, 64)
	if err != nil {
		return entity.Point{}, &InvalidRowError{Line: lineNo, Msg: fmt.Sprintf("y の数値変換に失敗: %v", err)}
	}
	var isA bool
	if series == "A" {
		isA = true
	} else if series == "B" {
		isA = false
	} else {
		return entity.Point{}, &InvalidRowError{Line: lineNo, Msg: fmt.Sprintf("series は A または B のいずれかである必要があります: %q", series)}
	}
	return entity.Point{X: x, Y: y, IsA: isA}, nil
}

// ParseRecords は CSV の全行を受け取り、ヘッダ行をスキップして全てのデータ行を解析する。
// 成功すれば []entity.Point を返し、途中でエラーが見つかればその行番号入りのエラーを返す。
func ParseRecords(records [][]string) ([]entity.Point, error) {
	if len(records) <= 1 {
		return []entity.Point{}, nil
	}
	pts := make([]entity.Point, 0, len(records)-1)
	for i, rec := range records {
		if i == 0 {
			// ヘッダ行をスキップ
			continue
		}
		lineNo := i + 1
		p, err := ParseRecord(lineNo, rec)
		if err != nil {
			return nil, err
		}
		pts = append(pts, p)
	}
	return pts, nil
}
