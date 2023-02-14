# frozen_string_literal: true

class Lasagna
  ##
  # Represents a recipe for lasagna.
  EXPECTED_MINUTES_IN_OVEN = 40
  def remaining_minutes_in_oven(actual_minutes_in_oven)
    EXPECTED_MINUTES_IN_OVEN - actual_minutes_in_oven
  end

  def preparation_time_in_minutes(layers)
    layer_prep_time = 2
    layer_prep_time * layers
  end

  def total_time_in_minutes(number_of_layers:, actual_minutes_in_oven:)
    time_preping = preparation_time_in_minutes(number_of_layers)
    time_preping + actual_minutes_in_oven
  end
end
