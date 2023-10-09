.global plus, main

plus:
    add x0, x0, x1
    ret

main:
    mov x1, #3
    mov x0, #4
    bl plus

    // プログラムの終了
    mov x8, #93  // システムコール番号: exit
    svc 0        // システムコールを発行
