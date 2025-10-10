package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/suzuki3jp/mn/internal/calc"
)

// WriteJSONToFile writes the given result as JSON to the specified file path.
// If the operation fails, it returns an error.
func WriteJSONToFile(result calc.MnResult, filePath string) error {
	// JSON に変換
	resultData := map[string]interface{}{
		"ro": result.Ro,
		"re": result.Re,
		"R":  result.R,
		"Z":  result.Z,
	}

	absZ := result.Z
	if absZ < 0 {
		absZ = -absZ
	}

	switch {
	case absZ >= 2.576:
		resultData["Z-jadge"] = "1%有意"
	case absZ >= 1.960:
		resultData["Z-jadge"] = "5%有意"
	default:
		resultData["Z-jadge"] = "有意差なし"
	}

	jsonData, err := json.MarshalIndent(resultData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal result to JSON: %w", err)
	}

	// ファイルを作成
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// JSON データを書き込む
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
	}

	return nil
}
