package com.crazyvinvin.aoc2024;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class InputReader {
    private Scanner getScannerOfInputFile() {
        File inputFile = new File("inputText.txt");

        Scanner scanner;
        try {
            scanner = new Scanner(inputFile);
            return scanner;
        } catch (FileNotFoundException e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
            throw new RuntimeException("Could not create the scanner");
        }
    }

    public Input getInput() {
        Scanner scanner = getScannerOfInputFile();
        Input input = new Input();
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            String[] valuesOnLine = line.split(" ");
            input.leftRow.add(Integer.parseInt(valuesOnLine[0].trim()));
            input.rightRow.add(Integer.parseInt(valuesOnLine[1].trim()));
        }

        return input;
    }
}
