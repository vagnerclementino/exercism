local Hamming = {}

function Hamming.compute(a,b)
    local diff = 0
    if #a ~= #b then
        diff = -1
    else
        for i=1, #a do
           if string.sub(a, i,i) ~= string.sub(b,i,i) then
               diff = diff + 1
           end
        end
    end
    return diff
end

return Hamming
