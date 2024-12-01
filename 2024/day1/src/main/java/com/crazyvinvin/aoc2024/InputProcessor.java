package com.crazyvinvin.aoc2024;

import java.util.ArrayList;
import java.util.List;

public class InputProcessor {
    private Input input;

    public InputProcessor(Input input) {
        this.input = input;
    }
    
    public int processInput() {
        List<Integer> difs = new ArrayList<>();

        while (!input.leftRow.isEmpty()) {
            int smallestLeft = getAndRemoveSmallestNumber(input.leftRow);
            int smallestRight = getAndRemoveSmallestNumber(input.rightRow);
            difs.add(calculateDifference(smallestLeft, smallestRight));
        }

        return handleAnswer(sumList(difs));
    }

    private int handleAnswer(int answer) {
        System.out.println(answer);
        return answer;
    }

    private int sumList(List<Integer> numberList) {
        int sum = 0;
        for (int i = 0; i < numberList.size(); i++) {
            sum+=numberList.get(i);
        }
        return sum;
    }

    private int calculateDifference(int x, int y) {
        if (x > y) return x - y;
        return y - x;
    }

    private int getAndRemoveSmallestNumber(List<Integer> integerList) {
        int size = integerList.size();
        int smallest = integerList.getFirst();
        int indexOfSmallest = 0;
        for (int i = 0; i < size; i++) {
            int val = integerList.get(i);
            if (val < smallest) {
                smallest = val;
                indexOfSmallest = i;
            }
        }

        integerList.remove(indexOfSmallest);

        return smallest;
    }
}
