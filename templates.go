package main



func MainTest() Request2 {
	var answers []Question

	answer := Question{Key: "entry.555539191"}
	AppendOption(&answer, "0-12", 2)
	AppendOption(&answer, "12-18", 8)
	AppendOption(&answer, "18-30", 8)
	AppendOption(&answer, "30-60", 8)
	answers = append(answers, answer)

	answer = Question{Key: "entry.1951481766"}
	AppendOption(&answer, "Vīrietis", 2)
	AppendOption(&answer, "Sieviete", 3)
	answers = append(answers, answer)

	answer = Question{Key: "entry.391036514"}
	AppendOption(&answer, "Jā", 2)
	answers = append(answers, answer)

	/*answer = Question{Key: "entry.34135510"}
	AppendOption(&answer, "Jā", 2)
	answers = append(answers, answer)*/

	answer = Question{Key: "fvv"}
	AppendOption(&answer, "1", 2)
	answers = append(answers, answer)

	answer = Question{Key: "fbzx"}
	AppendOption(&answer, "2753218881828164255", 1)
	answers = append(answers, answer)


	//https://docs.google.com/forms/d/e/1FAIpQLSeD9mbpMG4dgEFvbCOg_ra9HkF28hd_SUBtvu0v4H_VyrQ_NQ/formResponse"


	return Request2{
		Form: "https://docs.google.com/forms/d/e/1FAIpQLSeD9mbpMG4dgEFvbCOg_ra9HkF28hd_SUBtvu0v4H_VyrQ_NQ/formResponse",
		Questions: answers,
	}
}
