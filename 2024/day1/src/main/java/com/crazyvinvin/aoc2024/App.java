package com.crazyvinvin.aoc2024;


public class App 
{
    public static void main( String[] args )
    {
        InputReader inputReader = new InputReader("day1/");
        Input input = inputReader.getInput();

        IInputProcessor inputProcessor;

        // inputProcessor = new InputProcessorPartOne(input);
        inputProcessor = new InputProcessorPartTwo(input);
        
        inputProcessor.processInput();
    }
}