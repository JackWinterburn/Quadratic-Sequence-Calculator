defmodule M do
    def get_sequence do
        sequence = IO.gets("what is the sequence? ")
                    |> String.trim
                    |> String.split
                    |> Enum.map(&String.to_integer/1)
        IO.inspect sequence, label: "The Sequence is"
        gap1 = []
        calc_steps(sequence, [], 0)
            |> Enum.at(0)
            |> (fn (val) -> [gap1 | val] end).()
        IO.inspect gap1, label: "gap1"
    end

    # Calculate the steps in the sequence provided
    def calc_steps(seq, steps, n) when n > length(seq) do
        unless n == length(seq) - 1 do
            step = Enum.at(seq, n+1) - Enum.at(seq, n) |> IO.puts
        else 
            steps
        end
    end

    def calc_steps(seq, steps, n) do
        unless n == length(seq) - 1 do
            step = Enum.at(seq, n+1) - Enum.at(seq, n) |> IO.puts
            calc_steps(seq, steps ++ [step], n+1)
        else
            steps
        end
    end
end
