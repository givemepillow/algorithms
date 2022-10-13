s1 = 'fish'
s2 = 'fosh'


max_len = max(len(s1), len(s2))

matrix = {i: {j: 0 for j in range(len(s2))} for i in range(len(s1))}
max_sub_len = 0
for i, row in matrix.items():
    for j in row:
        if s1[i] == s2[j]:
            if i > 0 and j > 0:
                matrix[i][j] = 1 + matrix[i - 1][j - 1]
            else:
                matrix[i][j] = 1
            max_sub_len = max(max_sub_len, matrix[i][j])
print(max_sub_len)
