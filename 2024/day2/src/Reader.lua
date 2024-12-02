local Reader = {}

function Reader.readInput()
    local reports = {};
    for report in io.lines("../../inputText.txt") do
        local levels = {};
        for level in report:gmatch("%w+") do
            table.insert(levels, level);
        end
        table.insert(reports, levels);
    end
    return reports;
end

return Reader