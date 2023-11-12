package telegram

const(
	txtPrice = `Стоимость курса 3000руб`

	txtStart = `Добро пожаловать в мир Женского ДАО
	 - это мудрость и знание целителей и  
	   практиков Древнего Китая
	 - мягкая и глубокая работа с 
	   женским телом,энергей и 
	   сознанием - для комплексного 
	   оздоровления, молодости и  
	   красоты женщины, развития и 
	   обогощения её энергетики, для 
	   гармоничных женских состояний`
	txtMenu = `
	 /menu - стартовое меню
	 /price - стоимость курса
	 /bankdetails - рекизтты для  оплаты 
	 /payfor - оплатил за курс
	 /startvideo  - пробный урок`
	txtStartMenuChatAdmin = `/menu - стартовое меню
	 /delete - удалить юзера
	 /ignore - игнорировать
	 /accept - удалить`
	txtBankDetails = `Оплата по спб на телефон  +7978________ банк рнкб Ирина Федорищева`
	txtVideo       = `Пробный урок`
	txtPayFor      = `Дожидайтесь после оплаты Вам придет сообщение с ссылкой на закрытый канал`
	txtDeleteUser       = `Введите имя user для удаления из канала`
)
var btnStart = []types.TgRowButtons{
	{types.TgInlineButton{DisplayName: "Оплата", Value: "/add_cat"}, types.TgInlineButton{DisplayName: "Удалить", Value: "/add_rec"}},
	{types.TgInlineButton{DisplayName: "Отчет за неделю", Value: "/report_w"}, types.TgInlineButton{DisplayName: "Отчет за месяц", Value: "/report_m"}, types.TgInlineButton{DisplayName: "Отчет за год", Value: "/report_y"}},
	{types.TgInlineButton{DisplayName: "Ввести данные за прошлый период", Value: "/add_tbl"}},
	{types.TgInlineButton{DisplayName: "Выбрать валюту", Value: "/choice_currency"}, types.TgInlineButton{DisplayName: "Установить бюджет", Value: "/set_limit"}},
}