default rel

section .text
global nucleotide_counts
nucleotide_counts:
    mov qword [rsi], 0
    mov qword [rsi + 8], 0
    mov qword [rsi + 16], 0
    mov qword [rsi + 24], 0

check_chars:
    mov cl, [rdi]
    cmp cl, 0
    je done

check_A:
    cmp cl, 'A'
    jne check_C
    inc qword [rsi]
    inc rdi
    jmp check_chars

check_C:
    cmp cl, 'C'
    jne check_G
    inc qword [rsi + 8]
    inc rdi
    jmp check_chars

check_G:
    cmp cl, 'G'
    jne check_T
    inc qword [rsi + 16]
    inc rdi
    jmp check_chars

check_T:
    cmp cl, 'T'
    jne invalid
    inc qword [rsi + 24]
    inc rdi
    jmp check_chars

invalid:
    mov qword [rsi], -1
    mov qword [rsi + 8], -1
    mov qword [rsi + 16], -1
    mov qword [rsi + 24], -1

done:
    ret
