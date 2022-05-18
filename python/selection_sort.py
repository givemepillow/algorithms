def sort(array):
    new_array = []
    while len(array):
        new_array.append(array.pop(array.index(min(array))))
    return new_array


some_array = [10, 1, 2, 0, 2, 3, 4, 6, 4, 2, 3, 4545, 6, 7, 43]
print(*sort(some_array))
