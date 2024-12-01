package com.crazyvinvin.aoc2024;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class InputProcessorPartTwo implements IInputProcessor {
    private Input input;

    public InputProcessorPartTwo(Input input) {
        this.input = input;
    }

    public int processInput() {
        Map<Integer, Integer> summedDuplicates = getSummedDuplicatesMap(input.leftRow);
        int similarityScore = calculateSimilarityScore(summedDuplicates, input.rightRow);
        return handleAnswer(similarityScore);
    }

    private int handleAnswer(int answer) {
        System.out.println("Answer: " + answer);
        return answer;
    }

    public static int calculateSimilarityScore(Map<Integer, Integer> summedDuplicates, List<Integer> integers) {
        Integer similarityScore = 0;
        for (Integer key : summedDuplicates.keySet()) {
            int foundNumberOfTimes = containsNumberOfTimes(integers, key);
            similarityScore = (foundNumberOfTimes * summedDuplicates.get(key)) + similarityScore;
        }
        return similarityScore;
    }

    public static int containsNumberOfTimes(List<Integer> integers, int x) {
        int size = integers.size();
        int foundNumberOfTimes = 0;
        for (int i = 0; i < size; i++) {
            if (integers.get(i) == x) {
                foundNumberOfTimes++;
            }
        }
        return foundNumberOfTimes;
    }


    public static Map<Integer, Integer> getSummedDuplicatesMap(List<Integer> integers) {
        Map<Integer, Integer> summedDuplicates = new HashMap<>();

        int size = integers.size();
        for (int i = 0; i < size; i++) {
            int num = integers.get(i);
            if (summedDuplicates.containsKey(num)) {
                summedDuplicates.put(num, summedDuplicates.get(num) + num);
            } else {
                summedDuplicates.put(num, num);
            }
        }

        return summedDuplicates;
    }
}
