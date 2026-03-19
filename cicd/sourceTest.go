package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Простой тест: проверяем, что наша программа собирается
	fmt.Println("Running test: build test")
	
	// Переходим в папку src и запускаем make
	cmd := exec.Command("make", "-C", "../src", "all")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		fmt.Printf("❌ Build failed: %v\n%s\n", err, output)
		return
	}
	
	fmt.Println("✅ Build successful")
	
	// Проверяем, что бинарник создался
	cmd = exec.Command("ls", "../src/matrix")
	err = cmd.Run()
	
	if err != nil {
		fmt.Println("❌ Binary not found")
		return
	}
	
	fmt.Println("✅ Binary exists")
	fmt.Println("All tests passed!")
}
