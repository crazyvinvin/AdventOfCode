package com.crazyvinvin.aoc2024;

import static org.junit.Assert.assertSame;
import static org.junit.Assert.assertTrue;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Test;

public class AppTest 
{
    @Test
    public void shouldBeCorrect()
    {
        Input input = new Input();
        input.leftRow.add(1);
        input.leftRow.add(2);
        input.leftRow.add(3);
        input.rightRow.add(3);
        input.rightRow.add(4);
        input.rightRow.add(6);

        int answer = new InputProcessorPartOne(input).processInput();

        assertSame(answer, 7);
    }

    @Test
    public void summedDuplicatesListGeneration() {
        List<Integer> integers = new ArrayList<>();
        integers.add(3);
        integers.add(3); // 6/3
        integers.add(4);
        integers.add(4);
        integers.add(4); // 12/4
        integers.add(7); // 7/7
        integers.add(6);
        integers.add(6);
        integers.add(6); // 18/6

        Map<Integer, Integer> expectedSummedDuplicates = new HashMap<>();
        expectedSummedDuplicates.put(3, 6);
        expectedSummedDuplicates.put(4, 12);
        expectedSummedDuplicates.put(7, 7);
        expectedSummedDuplicates.put(6, 18);

        Map<Integer, Integer> summedDuplicates = InputProcessorPartTwo.getSummedDuplicatesMap(integers);

        assertTrue(expectedSummedDuplicates.equals(summedDuplicates));
    }
}
