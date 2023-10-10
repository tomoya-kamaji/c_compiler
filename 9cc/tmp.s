Token: Kind=TK_NUM, Value=12
Token: Kind=TK_RESERVED, Symbol=+
Token: Kind=TK_NUM, Value=34
Token: Kind=TK_RESERVED, Symbol=-
Token: Kind=TK_NUM, Value=5
Token: Kind=TK_EOF
.global main
main:
  mov x0, 12
  add x0, x0, 34
  sub x0, x0, 5
  ret
