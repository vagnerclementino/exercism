defmodule WordCount do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    String.replace(sentence, "_", " ")
    |> String.replace(",", "")
    |> String.replace(":", "")
    |> String.replace("!", "")
    |> String.replace("&", "")
    |> String.replace("@", "")
    |> String.replace("$", "")
    |> String.replace("%", "")
    |> String.replace("^", "")
    |> String.downcase()
    |> String.split(" ", trim: true)
    |> Enum.reduce(%{}, fn word, result ->
      Map.update(result, word, 1, fn counter -> counter + 1 end)
    end)
  end

  def normalize(word) do
    word
    |> String.replace(word, ",", "")
  end
end
