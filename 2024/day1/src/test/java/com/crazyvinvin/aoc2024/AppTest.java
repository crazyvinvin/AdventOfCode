package com.crazyvinvin.aoc2024;

import static org.junit.Assert.assertSame;

import org.junit.Test;

public class AppTest 
{
    @Test
    public void shouldAnswerWithTrue()
    {
        Input input = new Input();
        input.leftRow.add(1);
        input.leftRow.add(2);
        input.leftRow.add(3);
        input.rightRow.add(3);
        input.rightRow.add(4);
        input.rightRow.add(6);

        int answer = new InputProcessor(input).processInput();

        assertSame(answer, 7);
    }
}
