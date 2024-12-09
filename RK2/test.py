import unittest

# Импортируем классы и функции из вашего модуля
from main import Part, Manufacturer, ManufacturerPart, \
    query_manufacturers_starting_with_a, \
    query_manufacturers_with_max_part_price, \
    query_all_related_parts_and_manufacturers

class TestManufacturersAndParts(unittest.TestCase):

    def setUp(self):
        # Настройка данных для тестов
        self.manufacturers = [
            Manufacturer(1, "АвтоКомплект"),
            Manufacturer(2, "АльфаТех"),
            Manufacturer(3, "БетаМаш")
        ]

        self.parts = [
            Part(1, "Колесо", 500, 1),
            Part(2, "Двигатель", 5000, 1),
            Part(3, "Фара", 300, 2),
            Part(4, "Коробка передач", 2500, 2),
            Part(5, "Шестерня", 100, 3)
        ]

        self.manuf_parts = [
            ManufacturerPart(1, 1),
            ManufacturerPart(2, 1),
            ManufacturerPart(3, 2),
            ManufacturerPart(4, 2),
            ManufacturerPart(5, 3),
        ]

    def test_query_manufacturers_starting_with_a(self):
        expected_result = {
            "АвтоКомплект": ["Колесо", "Двигатель"],
            "АльфаТех": ["Фара", "Коробка передач"]
        }
        result = query_manufacturers_starting_with_a()
        self.assertEqual(result, expected_result)

    def test_query_manufacturers_with_max_part_price(self):
        expected_result = {
            "АвтоКомплект": ("Двигатель", 5000),
            "АльфаТех": ("Коробка передач", 2500),
            "БетаМаш": ("Шестерня", 100)
        }
        result = query_manufacturers_with_max_part_price()
        self.assertEqual(result, expected_result)

    def test_query_all_related_parts_and_manufacturers(self):
        expected_result = {
            "АвтоКомплект": [("Колесо", 1, "АвтоКомплект"), ("Двигатель", 2, "АвтоКомплект")],
            "АльфаТех": [("Фара", 3, "АльфаТех"), ("Коробка передач", 4, "АльфаТех")],
            "БетаМаш": [("Шестерня", 5, "БетаМаш")]
        }
        result = query_all_related_parts_and_manufacturers()
        self.assertEqual(result["АвтоКомплект"], expected_result["АвтоКомплект"])
        self.assertEqual(result["АльфаТех"], expected_result["АльфаТех"])
        self.assertEqual(result["БетаМаш"], expected_result["БетаМаш"])


if __name__ == "__main__":
    unittest.main()