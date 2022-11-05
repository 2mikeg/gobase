export FIRST_VAR="test"
export SECOND_VAR=1
export THIRD_VAR="false"
export FOURTH_VAR=3.14

TEST_DIR="./tests/base_model_test.go"

go test -v $TEST_DIR