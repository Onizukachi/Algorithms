# Contains Duplicate

- Ссылка: https://leetcode.com/problems/contains-duplicate/description/
- Платформа: LeetCode
- Сложность: Easy

## Вводные
Дан массив целых чисел `nums`.
Нужно определить, есть ли в массиве хотя бы одно значение, которое встречается минимум дважды.

## Что ожидается по условию
- Вход:
  - `nums` — массив целых чисел
- Выход:
  - `true`, если существует повторяющийся элемент
  - `false`, если все элементы уникальны

## Ограничения
- `1 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

## Примеры
1. Input: `nums = [1,2,3,1]`  
   Output: `true`

2. Input: `nums = [1,2,3,4]`  
   Output: `false`

3. Input: `nums = [1,1,1,3,3,4,3,2,4,2]`  
   Output: `true`
