def sort(array):
    for i in range(1, len(array)):
        j = i
        while j and array[j - 1] > array[j]:
            array[j], array[j - 1] = array[j - 1], array[j]
            j -= 1
    return array


some_array = [10, 1, 2, 0, 2, 3, 4, 6, 4, 2, 3, 4545, 6, 7, 43, -1, -17]
print(*sort(some_array))
