from typing import Optional


def search(sorted_array: list[float], target: float) -> Optional[int]:
    left, right = 0, len(sorted_array) - 1
    while left <= right:
        mid = (right + left) // 2
        if sorted_array[mid] == target:
            return mid
        elif sorted_array[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return None


assert search([], 0) is None
assert search([1], 1) == 0
assert search([1, 2], 3) is None
assert search([0, 1, 2, 3, 4, 5, 5], 5) == 5
assert search([0, 1, 2, 3, 4, 5, 6], 6) == 6, f"{search([0, 1, 2, 3, 4, 5, 6], 6)} != 6"
assert search([0, 1, 2, 3, 4, 5, 6], 3) == 3
print("TESTS DONE!")
