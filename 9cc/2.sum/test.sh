assert() {
  expected="$1"
  input="$2"

  # コンパイルして実行
  ./9cc "$input" > tmp.s
  cc -o tmp tmp.s
  ./tmp
  actual="$?"

  # 実行結果と期待結果を比較
  if [ "$actual" = "$expected" ]; then
    echo "OK: input:$input => expected:$actual"
  else
    echo "NG: $input => $expected expected, but got $actual"
    exit 1
  fi
}

assert 0 0
assert 42 42
assert 21 "5+20-4"
