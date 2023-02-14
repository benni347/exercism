# frozen_string_literal: true

class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    re = /^\[[A-Z]+\]:\s*[\\ntr]*(.*?)\s*(\\[rnt]*?(\\[rnt])*?)*?$/
    match = re.match(@line)
    match[1] if match
  end

  def log_level
    re = /^\[([A-Z]+)\]:\s*[\\ntr]*(.*?)\s*(\\[rnt]*?(\\[rnt])*?)*?$/
    match = re.match(@line)
    match[1].downcase if match
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
