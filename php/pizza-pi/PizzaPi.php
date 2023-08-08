<?php

class PizzaPi
{
    public function calculateDoughRequirement($pizza_amount, $person_amount)
    {
        return $pizza_amount * (($person_amount * 20) + 200);
    }

    public function calculateSauceRequirement($pizza_amount, $sauce_volume)
    {
        $sauce_per_pizza = 125;
        return ($pizza_amount * $sauce_per_pizza) / $sauce_volume;
    }

    public function calculateCheeseCubeCoverage($side_length, $thickness, $diameter)
    {
        return (int) (($side_length ** 3) / ($thickness * pi() * $diameter));
    }

    public function calculateLeftOverSlices($pizza_amount, $friend_amount)
    {
        $total_slices = $pizza_amount * 8;
        return $total_slices % $friend_amount;
    }
}
