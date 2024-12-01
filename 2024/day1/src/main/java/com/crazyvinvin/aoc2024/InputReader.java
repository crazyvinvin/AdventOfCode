package com.crazyvinvin.aoc2024;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class InputReader {
    private static String fileName = "inputText.txt";

    private String pathToInputFile;

    public InputReader(String pathToInputFile) {
        this.pathToInputFile = pathToInputFile;
    }

    private Scanner getScannerOfInputFile() {
        String pathName = pathToInputFile + fileName;
        File inputFile = new File(pathName);
        System.out.println("InputReader is reading from: " + inputFile.getPath());

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
            String[] valuesOnLine = line.split("   ");
            input.leftRow.add(Integer.parseInt(valuesOnLine[0].trim()));
            input.rightRow.add(Integer.parseInt(valuesOnLine[1].trim()));
        }

        return input;
    }
}
