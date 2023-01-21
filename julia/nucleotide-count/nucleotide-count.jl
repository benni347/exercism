"""
    count_nucleotides(strand)

The count of each nucleotide within `strand` as a dictionary.

Invalid strands raise a `DomainError`.

"""
function count_nucleotides(strand::AbstractString)
  counts = Dict((n, 0) for n in "ACGT")
  for c in strand
    c ∈ keys(counts) ||
      throw(DomainError("Nucleotide [$c] must be A,C,G or T"))
    counts[c] += 1
  end
  counts
end
