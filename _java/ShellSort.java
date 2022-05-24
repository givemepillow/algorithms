package _java;

public class ShellSort {
    public static int[] sort(int[] array, int l, int r) {
        int h = 1;
        while (h <= (r - l) / 9) h = 3 * h + 1;
        while (h > 0) {
            // System.out.println(h);
            for (int i = l + h; i <= r; ++i) {
                int j = i;
                int value = array[i];
                while (j >= l + h && value < array[j - h]) {
                    array[j] = array[j - h];
                    j -= h;
                }
                array[j] = value;
                // for (int v : array) System.out.printf(" %d", v);
                // System.out.println();
            }
            h /= 3;
        }
        return array;
    }

    public static void main(String[] args) {
        int[] array = {1, 4, 5, 6, 1, 2, 3, -2, 45, -90};
        for (int value : sort(array, 0, array.length - 1)) {
            System.out.printf(" %d", value);
        }
    }
}
