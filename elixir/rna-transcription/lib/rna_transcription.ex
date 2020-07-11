defmodule RnaTranscription do
  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> RnaTranscription.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    dna
    |> List.to_string
    |> String.codepoints
    |> Enum.map(fn(n) -> transpile(n) end)
  end
  
  def transpile(nucloide) do
    case nucloide do
       "G" -> ?C
       "C" -> ?G
       "T" -> ?A
       "A" -> ?U
    end
    
  end
end