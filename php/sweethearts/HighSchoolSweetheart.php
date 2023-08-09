<?php

class HighSchoolSweetheart
{
    public function firstLetter(string $name): string
    {
        return $name[0];
    }

    public function initial(string $name): string
    {
        return strtoupper($this->firstLetter($name) . ".");
    }

    public function initials(string $name): string
    {
        // Split the full name into first and last names
        list($firstName, $lastName) = explode(' ', $name);

        // Get the initials of the first and last names
        $firstInitial = $this->initial($firstName);
        $lastInitial = $this->initial($lastName);

        // Return the concatenated initials
        return $firstInitial . " " . $lastInitial;
    }

    public function pair(string $sweetheart_a, string $sweetheart_b): string
    {
        $first_initals = $this->initials($sweetheart_a);
        $second_initals = $this->initials($sweetheart_b);

        $heart = "     ******       ******
   **      **   **      **
 **         ** **         **
**            *            **
**                         **
**     $first_initals  +  $second_initals     **
 **                       **
   **                   **
     **               **
       **           **
         **       **
           **   **
             ***
              *";
        return $heart;
    }
}
