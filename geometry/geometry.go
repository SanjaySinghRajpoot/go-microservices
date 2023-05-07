package geometry

func Cal(len, bre, float64) (area, perimeter float64)
{
   area = (len * bre);
   perimeter = 2*(len + bre);

   return area, perimeter;
}