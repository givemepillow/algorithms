import typing


def sort(arr: typing.List[float | int]):
    match len(arr):
        case 0 | 1:
            return arr
        case 2:
            if arr[0] > arr[1]:
                arr[0], arr[1] = arr[1], arr[0]
            return arr
        case _:
            return [*sort([v for v in arr[1:] if v < arr[0]]), arr[0], *sort([v for v in arr[1:] if v >= arr[0]])]


print(sort([99, 1, 2, 3, 4, 5, 6, 7, 8, 9]))
