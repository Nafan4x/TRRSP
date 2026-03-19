#!/bin/bash

echo "=================================="
echo "Starting all tests"
echo "=================================="

# Делаем тестовые скрипты исполняемыми
chmod +x test2.sh

# Запускаем Makefile тест
echo ""
echo "--- Running Makefile test ---"
make test

# Запускаем shell тест
echo ""
echo "--- Running shell test ---"
./test2.sh

echo ""
echo "=================================="
echo "All tests completed"
echo "=================================="
