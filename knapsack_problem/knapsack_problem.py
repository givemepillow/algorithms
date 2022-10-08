from tabulate import tabulate

N = 4
G = 11

weights = [2, 4, 5, 7]
prices = [30, 60, 70, 80]

matrix = [[0 for _ in range(G + 1)] for _ in range(N + 1)]

# Для отслеживания выбранных предметов.
items = [[[] for _ in range(G + 1)] for _ in range(N + 1)]
coords = (0, 0)
max_price = 0
for i in range(1, 1 + N):
    for g in range(1 + G):
        prev = matrix[i - 1][g]
        if weights[i - 1] <= g:
            current = prices[i - 1] + matrix[i - 1][g - weights[i - 1]]
            if current > prev:
                # Отслеживание выбранных предметов.
                items[i][g] += [i] + items[i - 1][g - weights[i - 1]]
            matrix[i][g] = max(current, prev)
        else:
            matrix[i][g] = prev
            # Отслеживание выбранных предметов.
            items[i][g] = items[i - 1][g]

        # Отслеживание выбранных предметов.
        if matrix[i][g] > max_price:
            coords = (i, g)
            max_price = matrix[i][g]

table = tabulate(matrix[1:], [f"G{i}" for i in range(1 + G)], tablefmt="fancy_grid")
print(table)

print(f"Лучший вариант загрузки - {max(v for row in matrix for v in row)}.")
print(f"Предметы для загрузки: {items[coords[0]][coords[1]]}")
