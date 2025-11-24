package main

import (
 "bufio"
 "errors"
 "fmt"
 "os"
 "strconv"
 "strings"
)

type Matrix struct {
 Rows int
 Cols int
 Data [][]float64
}

func NewMatrix(rows, cols int) Matrix {
 data := make([][]float64, rows)
 for i := range data {
  data[i] = make([]float64, cols)
 }
 return Matrix{Rows: rows, Cols: cols, Data: data}
}

func (m Matrix) PrintMatrix() {
 for i := 0; i < m.Rows; i++ {
  for j := 0; j < m.Cols; j++ {
   fmt.Printf("%8.2f ", m.Data[i][j])
  }
  fmt.Println()
 }
}

func Add(m1, m2 Matrix) (Matrix, error) {
 if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
  return NewMatrix(0, 0), errors.New("размеры матриц не совпадают для сложения")
 }

 result := NewMatrix(m1.Rows, m1.Cols)
 for i := 0; i < m1.Rows; i++ {
  for j := 0; j < m1.Cols; j++ {
   result.Data[i][j] = m1.Data[i][j] + m2.Data[i][j]
  }
 }
 return result, nil
}

func ScalarMultiply(m Matrix, scalar float64) Matrix {
 result := NewMatrix(m.Rows, m.Cols)
 for i := 0; i < m.Rows; i++ {
  for j := 0; j < m.Cols; j++ {
   result.Data[i][j] = m.Data[i][j] * scalar
  }
 }
 return result
}

func Multiply(m1, m2 Matrix) (Matrix, error) {
 if m1.Cols != m2.Rows {
  return NewMatrix(0, 0), errors.New("количество столбцов первой матрицы должно быть равно количеству строк второй матрицы для умножения")
 }

 result := NewMatrix(m1.Rows, m2.Cols)
 for i := 0; i < m1.Rows; i++ {
  for j := 0; j < m2.Cols; j++ {
   sum := 0.0
   for k := 0; k < m1.Cols; k++ {
    sum += m1.Data[i][k] * m2.Data[k][j]
   }
   result.Data[i][j] = sum
  }
 }
 return result, nil
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func readInt(prompt string) (int, error) {
 fmt.Print(prompt)
 input, _ := reader.ReadString('\n')
 input = strings.TrimSpace(input)
 val, err := strconv.Atoi(input)
 if err != nil {
  return 0, fmt.Errorf("неверный ввод: %w", err)
 }
 return val, nil
}

func readFloat64(prompt string) (float64, error) {
 fmt.Print(prompt)
 input, _ := reader.ReadString('\n')
 input = strings.TrimSpace(input)
 val, err := strconv.ParseFloat(input, 64)
 if err != nil {
  return 0.0, fmt.Errorf("неверный ввод: %w", err)
 }
 return val, nil
}

func getMatrixSize() int {
 for {
  size, err := readInt("Введите размер матрицы (2 для 2x2, 3 для 3x3): ")
  if err == nil && (size == 2 || size == 3) {
   return size
  }
  fmt.Println("Ошибка: Введите 2 или 3.")
 }
}

func readMatrixValues(rows, cols int) (Matrix, error) {
 fmt.Printf("Введите элементы матрицы %dx%d:\n", rows, cols)
