.global main

main:
    mov x0, #5       // x0に5をロード
    add x0, x0, #20  // x0に20を加算
    sub x0, x0, #4   // x0から4を減算
    ret              // 関数を終了
    