package.path = package.path .. ";day2/src/?.lua"

local reader = require("Reader")

local reports = reader.readInput();