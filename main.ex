defmodule M do
    def getSequence do
        sequence = String.split(IO.gets "what is the sequence? ")
        seq_length = length(sequence)
        putSequence(sequence, 0)
    end

    # print the sequence on the page recursively
    def putSequence(seq, n) when n > length(seq) do 
        IO.puts(Enum.at(seq, n))
    end

    def putSequence(seq, n) do
        IO.puts(Enum.at(seq, n))
        putSequence(seq, n+1)
    end
end
