.section .bss
buffer: .space 10    // 数値を文字列に変換するためのバッファ

.section .text
.global plus, main

plus:
    add x0, x0, x1
    ret

main:
    mov x1, #3
    mov x0, #8
    bl plus          // plus関数の呼び出し
    bl int_to_str    // 結果を文字列に変換

    // 結果を標準出力に書き出す
    mov x8, #64      // システムコール番号: write
    mov x0, #1       // ファイルディスクリプタ: 標準出力
    ldr x1, =buffer  // データのアドレス
    ldr x2, =10      // バイト数の上限
    svc 0            // システムコールを発行

    // プログラムの終了
    mov x8, #93      // システムコール番号: exit
    mov x0, #0       // 終了ステータス: 0
    svc 0            // システムコールを発行

int_to_str:
    // 簡単な整数を文字列に変換する関数
    // 入力: x0 - 整数
    // 出力: bufferに文字列としての整数
    // 注意: この関数は非常に単純で、10未満の正の整数のみをサポートします
    ldr x1, =buffer
    add x0, x0, #48    // ASCII変換
    strb w0, [x1]
    mov w0, #10        // 改行
    strb w0, [x1, #1]
    ret
