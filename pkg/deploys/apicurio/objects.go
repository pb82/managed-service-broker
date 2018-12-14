package apicurio

import (
	brokerapi "github.com/integr8ly/managed-service-broker/pkg/broker"
)

// Apicurio plan
func getCatalogServicesObj() []*brokerapi.Service {
	return []*brokerapi.Service{
		{
			Name:        "apicurio",
			ID:          "apicurio-service-id",
			Description: "apicurio",
			Metadata: map[string]string{
				"serviceName": "apicurio",
				"serviceType": "apicurio",
				"imageUrl":    "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/4gKgSUNDX1BST0ZJTEUAAQEAAAKQbGNtcwQwAABtbnRyUkdCIFhZWiAH4gABABAADQAZAB9hY3NwQVBQTAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA9tYAAQAAAADTLWxjbXMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAtkZXNjAAABCAAAADhjcHJ0AAABQAAAAE53dHB0AAABkAAAABRjaGFkAAABpAAAACxyWFlaAAAB0AAAABRiWFlaAAAB5AAAABRnWFlaAAAB+AAAABRyVFJDAAACDAAAACBnVFJDAAACLAAAACBiVFJDAAACTAAAACBjaHJtAAACbAAAACRtbHVjAAAAAAAAAAEAAAAMZW5VUwAAABwAAAAcAHMAUgBHAEIAIABiAHUAaQBsAHQALQBpAG4AAG1sdWMAAAAAAAAAAQAAAAxlblVTAAAAMgAAABwATgBvACAAYwBvAHAAeQByAGkAZwBoAHQALAAgAHUAcwBlACAAZgByAGUAZQBsAHkAAAAAWFlaIAAAAAAAAPbWAAEAAAAA0y1zZjMyAAAAAAABDEoAAAXj///zKgAAB5sAAP2H///7ov///aMAAAPYAADAlFhZWiAAAAAAAABvlAAAOO4AAAOQWFlaIAAAAAAAACSdAAAPgwAAtr5YWVogAAAAAAAAYqUAALeQAAAY3nBhcmEAAAAAAAMAAAACZmYAAPKnAAANWQAAE9AAAApbcGFyYQAAAAAAAwAAAAJmZgAA8qcAAA1ZAAAT0AAACltwYXJhAAAAAAADAAAAAmZmAADypwAADVkAABPQAAAKW2Nocm0AAAAAAAMAAAAAo9cAAFR7AABMzQAAmZoAACZmAAAPXP/bAEMABQMEBAQDBQQEBAUFBQYHDAgHBwcHDwsLCQwRDxISEQ8RERMWHBcTFBoVEREYIRgaHR0fHx8TFyIkIh4kHB4fHv/bAEMBBQUFBwYHDggIDh4UERQeHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHh4eHv/CABEIAZABkAMBIgACEQEDEQH/xAAcAAEAAgMBAQEAAAAAAAAAAAAABgcEBQgCAwH/xAAbAQEAAgMBAQAAAAAAAAAAAAAABAUCAwYHAf/aAAwDAQACEAMQAAABuUAAAAAAAAA+Z9EXjBZ6mMIvRRX0LxVBviwWo24AAAAAAAAAAAAAAAAAflaFlQOjtEWNBcMAAAAfu80QuCzeUt2dYoLOgAAAAAAAAAAAAABHoxQhJIkAAAAAAlpFJ5ccrIXM/wBABqo7rlTdV+4xkTgbq0AAAAAAAAABUH0o8fgAAAAAPpldEkVt30Bp/me41NWw6H0VjwvVIvQBrmAWZZ/MlnzeXswTuWAAAAAAAAQKUcrmF5AAAAABn4/SxkScHjDpbTZSmt/mre1DCUAAA/fwXNNebL/seL2glUQAAAAAAiBUtefv4AAAAACwCeWgDUfSiI9v61hWdwDIbZhqVnzKTT0XILt/d9VVGfZDZDr38sN911ptpqA3VwAAAAADm6+OTjyAAAAADI6sqa8R8PvUGqbH9GVXoIyPmWPvZ7YEvnIhL/1O5geMtPto9LrlzZW+NhJtFVn3fbMQLb5R5MxsnZDD78AAAAAqGkZrCgAAAAB9PnPS+NuEWqS3ZLFvKu2k/Zxq6l+2Z6BC/uqXwyqMOF1Ewi+OjXYY7wAAPVo1ZINsC+xbeegAAAPHvTnK2J+/gAAAAAv+getTbAAAAY2S+feftLe1E1fehpswAAAHvwOlftpN3c+ZBlrAAARGXQU5tAAAAABtOuOW+pAAAAABz30JTMS+hQr+zAAAAAu+WQqa23nYbYQAACBzyBnN4AAAAAJp0xzh0eAAAAAKfuCn41zBBW9wAAAABcs2hM2tvPA2wQAAEDnkDObwAAAAAS3p/jbqkkAAAAAFIW9z1C6T5CD1oAAAAFyzaEza288DbBAAAQOeQM5vAAAAAAl0RHYv35i6AN6AAB4wKg02KKlX3YfNoAAAAFyzaEza288DbBAAAQOeQM5vAAAAAAA9+BNJXUH6dk+tJsPn2vI3F/yq9B+nzNU8AAAAAAC5ZtCZtbeeBtggAAIHPIGc3gAAAAAHo8yWz7XKYlU+AEUjtmtNhQkf6bgka8p978Q+lAAAAAAuWbQmbW3ngbYIAACBzyBnN4AAAAAFy1L1wZIAAAAK6qnpjnWv7DBEToQAAAALlm0Jm1t54G2CAAAgc8gZzeAAAAACbdLc69FAAAAACl7op+NdQQVvbgAAAAXLNoTNrbzwNsEAABA55Azm8AAAAAFhdFc69FAAAAACn7gp+NcwQVvcAAAAAXLNoTNrbzwNsEAABA55EjmAAAAAAFhdFUFfoAAAAAp+4Koj3Fdis7kAAAAC5ZtE5ZbedhthAAAMbJHHePaFXgAAAAzS7bSws0AAAAAQmbfHCRzUz8Cn9HBkAAA9+JplotvMLjzYPvwAAADW8tdbxw5WbrSgAA9nm/sC3AAAAAAACGUx01AIfR1G+nzgdcAANgx+d+4W/suJCTTAAAAAAYlSXMOT9L2VinH2X1lmHPNwSgAAAAAAAAAaStrkabDm7F6axo9zzfsugMj7jV1jZaTSBshgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAf/EACwQAAEEAgECBgICAgMAAAAAAAQCAwUGATAANTYQEhMUIEAHERUWIVAkYID/2gAIAQEAAQUC/wCiuLQ3gqwQ4/H7pHJ45eM8zdy+f3Y3ibuTxq7tcGt8Q7wWTjyv9MpSUpk7XGCckLfJkcIKIIVpCl5IPMfdXMcjJiOkcf6BWcJxM24QXknLHyKt0dGSJma6JOjY+9MzAcW1Nzx0pnbjGc5iarIm8jazFhcxjGMfes9mbA4S+6S9tg4E2UVDQAEZj4kSII/H7MAjjlqVyGsDZbn17bZfJzP+c7G0KcXXaklHEpwlPidJBhYNtC88KkTSs/Cuzv6+tc7B6O4IV8wmtwA8U34yMkIAmSsJZOc5znPzrEz9S3zeIwVWcqzsAEfNKr8MPEjeC1pQiYsfHFrcXqrEr7xr6Msc1HAHlvGl7GGnHnqzCtxIngcWwExMSxEgvYM84O/FmIOD+hd5X30hto8J7VjwlD2I8eRNfOI+bAhT/Ga9Jucbqz/E1Vvn9WG5mqj8XVcchIouNJ322S/jonbS4j+QO8JI1oEaQMeOJ+IMcYZwOr8Fio8bmMYx9e7SHvZjYOy4Q/DANxsfx91thmZkHJAr4RsUYdmOr4Y3MYxjH2JszAEUrOVK2fj2N9R/wtcn7h7xYZdfdiK401xOMJx4rUlGHpeOZ47Zo9PF2pvmbUvmLU5xNrxxq0B54zOxjnGSGHt35IK/TGxCVLXDhpAjeWSQ9iEFHGG5HqzmcYqwvHKqzyNjxwGvB11tpBtlDa4XYpB7jzzz2dCc5TmnHOuq2XUn3Ng2UYP3U54Zi0lHJxhOPjY5QsDhJL5K9lef9vL61q8qC3MvFbPx4J6UVpJYbJYmAFx5exCvItlfqM6px30YfbDD+1itVjBwZHbYJfqQ+q4r8lb2RTXryeyZY9tJ7KkrzQuq95/Vc2VRPnsWy5J/UtspvSdV87d2UlP7suy69S2UzpOq+du7KFj92LZdepbKZ0nVfO3dlQIwNYNlsew7MbKZ0nVfO3diVZSqvyCJOM1SRSAg3Vqcc2UzpOq+du7axMKiTR3m32dC1JQiwyn8gRtpnSdV87d3V6eJiXIqWBkm/kccKEiamXpDO6mdJ1Xzt3ehSkKEs8wOmMub3roVhaeSVmUl5+fk3cLWpat9M6Tqvnbv0oFtxqGNSpYmf8Z+jTOk6r527vi4OSkcDUhXIyqRwb/hIwIZjr9V4fEHB430zpOq+du7U4ypVZqrbKMYxjHznIBshK0qQvbTOk6r527toMPjOu3RuFt7aZ0nVfO3dgbKiShWUDj6lpwtB7GRjdlM6TqvnbuykNYdsey4N+WX2UzpOq+du7KB3BsuvUtlM6TqvnbuygdwbLr1LZTOk6r527soHcGy69S2UzpOq+du7KB3BsuvUtlM6TquDXq1zZ+OWvNL7Luj/k7Km35IXUU1h8YhpTL+v8dC+nHbLiP6sZrQnKlBM+3E13+N9A/UCM4WWCOgQPY82l1k8ZYheqpg+4P2SYbMgFKgvxxmmkQmQ2d1ni/esaQx3SiIwNsETbPRA8sLLRhcYR8kJUtVVrHoq+hYIL3HHEKbX8gQyDXoaMajmd5YzBbEtTOGRMkJnOP14MDEPqjalJE5hoICL+nJRYh+D64Yzl4d9nPgMAYRmPrCs8FGYFa+msYdfMBiY4lKU4+tnGM8yMPniWGE/wDjj//EADQRAAAEAwQGCAcBAAAAAAAAAAECAwQABRESIEBxBiExMjRREBMUMEFCYaEVIjNSYIGxI//aAAgBAwEBPwH8DSZOFdwgwpKXaZLZi4VNMyhrJAqMNNHjG1rjT0hCXt2+4Xpm8o2rohmGDYSxR4PIvOGrJJqWymF6dS7s5+tJuj7YGVywXZrRt0IIQpC2Shq6DqFTC0YaBC8+bJ6ifNCukSw7hQCDTp4Pm/kfFnn3wpNXKpBIcagOAYMzO1bAbPGEkypEAhNgQIgAVGHs/KT5UNfrC7hVcbSg16E2LlTdIMBJng+T3CBkr0PJ7hB5c6JtIMCUS6h79vMiMUbCIVMO0YPOHhh34WfuFy2FDVCG7c7hQEybYbaPJl1rDWEWqKP0y0uzJuC7YweOCARKNQiVPe1o1HeDbfWJYUMXlgtHlBK4EvML8xCjpTPBSHiwyG/M+LUzwWjrcROZYcr8z4tTPAsZYq7NyLzhBEiBATJsC/M+LUz78oVGkIyxqkAUJ3Mz4tTPAJTR0kFCnhlPwONhcKevcTPi1M8FIHgqJiibw/l+Z8WpngpCNHf6vzPi1M8FIeLDIb8z4tTPBSIaPA/d+YjV0pngmq/ULFU5QUwHKBguuVyt0hUN4QYwmGo4OTzUEv8AFXZ4DG3pOcpAtG2RNpn2o1gm6HvhWk0cNdRR1coJpJ9yfvCmkg+QkOn67r6g/iH/xAApEQABAgQFAwQDAAAAAAAAAAABAgMABBFAEiAhMTIQQVETFCIzMEJg/9oACAECAQE/Af4MuJTuYD6CaVtSQNTC5sfrCnVq3PWXf/VVm68G4W4pZ1zS7uIUNi+9g0G8E11PQAnaEyqzvAlE9zHtm/Eeg34gMIBqLB1z001gkk1PRuVJ1VCUBO3QuoHePcN+Y9y35gPIPewUyXVVVtAl2x2hLSEmoELUECphc2TxhS1K3OVlWFYs328Cs6TUVspsfGudngLKa+vOzwFlNr0w52eAsXXktwpRUanOzwFgp5Z7/hZ4CwUwhXaHJWmqfwM8BZTTdDiGdngLKa4Z2eAspr687PAWUz9edngLJacSaQRTKhOJVIGlnMMV+ScgFYYZwane1Wyhe8GT8GBJ+TCGko2/kP/EAEcQAAECAwIHDAcFBgcAAAAAAAECAwAEMBEhEiIjMUFRYRMyQEJScXJzgZGxwQUQFCBiodEkMzSS4UNQU2OCohVggIOTsvD/2gAIAQEABj8C/wAi4S1pSNZMHDnmydSMbwjJMTDnYBGT9H/mci6SY7zH4OX7zGNItHmUYysgsdFcWOF5npI+kZCcZXswr/3MVKIAGkwUNKMy4P4ebvgpl8CWT8N574wn33HT8SraX2ecdSOSTaO4wEz0sFfE3d8o+zTCSrkKuV+4SSbAILckPaXeVxB9YtmZhSk8gXJ7q4MpLOq+PMO+AJ+baca5BxlDt4fhTC7Vnetp3xgpWrcmNDSc3brrWC8wFvD2Vo6V5+6AS1u7nKdv+UWAWDh5lZOxyZ0nQj9YU8+4pxxWdRrYSE7kxpdVm7NcBTbe6PfxF5/097KzTY2W2xk0uu9lkZOUH9S43GZSllw706Dwhch6OXj5nHRo2CLTVCEJKlG4AQmZ9KDCVnDOgc8BKQABmA9zLvAHki8xZKMBPxLjLTC1bLbB7qZSdVsQ4fA8GV6OkV5TM6scXZWTLy7ZW4rRAcXY7NHOvVsHuZdzG0IGcwUsZBvZnPbFpNpoJkptextZ8OCbgwftTou+Aa4JJtJqolpdGG4qMFFi3lfeOa/09ZWtQSkZyYLMh2unygrWoqUc5NP2d85dAz8ocCcmncyRcNZ1Q5MvqtWs1UtNJKlqNgAi+xUyv7xXl6y8+uwaNZiw4jIzIHnVS80rBWk2iEPo/qGo8B9maVkGDZzq0msPSEyjLuDEB4qfXujpv4qdJguvK5hoFDIy7i+ZMWltDfSVGPNNjmFsY02rsTH4l3uEfine4RiTh7UQTu7bjK98OALUg2PO4jf1rbu8n7OwbT8R1esvunmGswXnjfoGoe9kGVEcrMIBm5j+lH1jJyyLdarzF3B1NpOSl8RPPpqoZaThLWbAIblUcUYx1nT6lOuqwUJFpMFxVzYuQnUPdtabsRy1ZoCnRu7nxZu6LALBwl+aOdKcXn0QVE2k1V+knBit4rfPp9fsbKsk2cY8o+4G2UFazoEB2dscXyOKPrFgFgHuWrUEjbGPNt9l8Ygec5kxiSajzri6TT+eL5NP54xpPuXGOy8juMfiMDpCyMk8hfMqtLyYO+OGqqEJFqlGwCGZVPETfz6fVYg5Zy5OzbGQZUoco3CMvMpTsQLYvmHvlGTmljnTGAym/jKOc+vDdWlCdZMFLAU+rXmEWIUllPwC+LXXVrPxG2jaCQYdlnnVLsGEnCNtV8DetWNjsz/OqlxQtQwMPt0eszk9jWXNtaANsWJAA94JZYuV+1N4jDfdU4dpqsK0E4J7ahUcwFsOunOtZVVXNEXvru5h/wCNJTLycJCs8FpV6TehWsVQoaDbCF8pINObc1NK8K0sxyGxTVYMq3jIrSyvgspzZ1gD5irLM8t1I+dV9nQFXc1VsclSh86b21SfGrJD+ZbVwuU2Kp60+VN3pp8asrswj/aarfVedVXWnypu9NPjVb2IV4VW+q86qutPlTd6afGrLKVvVHAPbVWB+zARVV1p8qbvTT41QpJsIvENzAOPvXBqVTcmF8UXbTCnFG1SjaaqutPlTd6afGthKtUwu5xPnCXmVhaFi0EUStZCUi8kxgt2hhG927ayutPlTd6afGvgjKy5OM2fKMKWeBVpQblD38KYdCdmkxuacmxydfPXV1p8qbvTT48ACkKKVDSIwfad1H8xNsBM+yjcznUjOmApJtBvB9Rbk20lI46tMWbsG+gLIwlqKjrPAFdafKm700+PA5Rt3fhoWw8lG+KCBFh4ErrT5U3emnx4BhMMHc+Wu4RbMzwGxCIDysOYWm9OHmHZ6y7jNLOfA0xkZv8AOmCpxrCRyk3jgCutPlTd6afGsEpFpOYQma9JIC3c4aOZPPFguFAvSYDb3J0KgoWClQuINZXWnypu9NPjW/xSYTsZB/7U/b2U4yfvNo11ldafKm700+NVqXTncWEw2w2LEoTgimUKFoIsMOsHiKsqq60+VN3pp8asvbmQFK+VXC5aAaqutPlTd6afGqOqVVb6rzqq60+VN3pp8ao6pVVvqvOqrrT5U3emnxqjqlVW+q86qutPlTd6afGqOqVVb6rzqq60+VObGpIV3Gq87oQz4mrLua0EVWzyyVU3GVZlpKYWyvfIUUmo7NEffLsHMKodAvaVb2VAkZzdDTI4iQKgn205N/fbFU25ZoWrcVYIalm962nBFVbSxalQsMOS686T30/aFDJs39uiq5KvjFWO464XLPpvGY8oa6Xt0yix9wYgPETX3dlOXb/uFJLDKbVKhLDejOdZrbm7iuJ3jgzpjcpluzkqGZXvhKElSjmAhM76STlBehrVtPATMygsd4yOVBQtJSoZwff3JhGEdJ0CLBjOq36+AFmZaS42dBguejXv9tz6xl5N1O2y0Rf6sFlhxw/Cm2AZiyVb+K9XdGEyjDd/iLz/AKcDyyLF6FjPBMvY+jZcYsdZWjnT68jLOK22QFTrlg5CPrG5MNhCdnBMdhpXOgRdKsf8YixIA5uD3i2L2GvyCMVlscyf9HH/xAAsEAEAAQEGBQQDAQADAAAAAAABEQAhMDFBUWFAgaGx8HGRwdEQIOFQYIDx/9oACAEBAAE/If8AguzU0FQInqKiMbqh71Mso8MCpNgN2ryH5oz7Uqd0De5SAlc5jurFv0E9jbQz/i2yNlICksgzB7PaaIkPMdV8FNlTNrkKMlHoB4MKfD83l7qEapsrk48v8ESwJVYAqK8rLUcz4UllCywL8DZGwEHNZUW45vovvPH2Xqzn0m9W7EtN8CQKWAZ1ETaorZt94qzlvJgomYLACw4+SuHx/X12UlAJbK3w60Fh6GarLIWlPJ9P1UCWsevRN7FSA9qRdanWbwyKkgCyeXbg8Rvxdduu9JRJXFvUvfByrRAeA3+n0osbQCA/TalbZ5FSxDK1fYpgnOXoCz9WtaYlh4Y8M5gXa+DTvrpTbeuuWwZbuhQkEW5Z2vr+kxCZNp5VMFtv3vpTtyYq2txN4iW/jhIwfJOulKWRKuK3qI0wGQauhRMgrUtWhp+RYBlEBTrZTAHZ80gLZRK3QoyMNbLUve9eCQfmjKFWzbnQMg2L0FrseVqykEfhbH5MEeDN0CnTzLd11N6giSKj6Fs92OBdZrhDZ5B731hBYNa2fq9vy4mOw8XzOp02R4OhcIR6nJRgNdP4mlx6S9M9/R+6IrazIj61eMHWo9SiNGck4BMsRYi48hSyy3uLwkuGT8miz8YHqw8dIqYarAw0j9kfXax5tH7oPyows3K71AQANDh2l/AsHN97OV6xwfXLUB7aavi/BjJIUrD4WZ/qC1XZf3R8FZiy7faiRgsALDiXUBfqdnU04JEq5t7vNnm8XI7/AJ2kCbPpP0wkuDQGdcepQaYEAEB+m7JFBU0ITJz6VMynoHVo3Qo+K+Uv5rNn0X1Wc/HakY3w+ysvrR/4oCfS5vlx+y2LDu3rugAZrhRvEhMzdq9/xK4522dJvmDvLQh1zuTQS0NEdHmsb/8Aa/r8tgjGEU2GMndam3LLuGlLlmty3YgMUxjG1QYJ2vZ0SPIvkb05MN9HVby/JOhxKzZ6lxoaVYAQH7EHh2WY6Rr61087cr1ZeOX7H1eFjZLlTvyi5t7YV2xy+90KSCA0vzhh5N62LiHKiBw9xLvUO9wJYKAciG+sS9bu0WFbucy+nNliuVnxd2hW+7le6do+ygggu3CoeIS+tad71Vv/AEHzdoAyb7b3bHoC/F6Jgxb1L1zt9niCWVovVdE7r3yenEEpXEy/TB83vRO698npxBJTgXysHWL1f5JTfF73vk9OIJTdHJo0ZqxsDH753bkGR9gqRkDbt75PTiSWFyxaZDcqTxRrEuT0jIQBTCidlz6r7yenFErTUsrDfQ0AhzuwfvsmWP0inoJtk7d/0v8AyenFksHpCQlBgBwC68an+vA03IzoLxwGCUoErAUlq4tXlNKZxjor3xpuy4rK8B5PTjiQGIOXEsrqFORQUBEsTgvJ6cWSAHL+W58qcCTfdV+KbdiCM2sMef5tK/LhHqGrGbfT5SlyjjZnrpwHk9OJJLAdAEq0kpZt852oEABAGVwE0Wps+tp1IwEI33k9OJJCW2FE9/gc7tXoOeCL7yenEEsEmPNigIDDsXc/hg1Gs5ZG5l0vfJ6cQSUDPslR1S9JyRL+uHwXvk9OIJed2veid175PTiCXndr3onde+T04gl53a96J3Xvk9OIJeV2veid174PS7RQWk8hek3qZgD2Pq9ciWTPo/29RrHq0fF3j/c5kUHkCG4xeLHz4LvN6kjDXqsfi8BuVAb0X8ct5ZUWEcB+z5u4PQbW9HnB9yM72bEZbNEzgh0ZN2xd0sr5vZSmMzyhuVCWfYyBdZDzq1Plv0GUGBjo+ulIiiQlzK40bBq7VaPYs/Ob6DVuSWnybUgdLVFqP7qUGASrUZLAvB93bgSZ2LgG83pIQwCE/civkDqtWfTYGOxocBj7FPrs0skxjP2+9KgoydYWUigQ7/g8qZLQy849IPmjCwsNty/Tg8JiRYz90uWW+DSxEax+AVgtocSWye7TQ+sS86ExfLNu68I3Lu8paQOp9FQAugjhwYAaJSUqd6HZZ2T/AKcf/9oADAMBAAIAAwAAABDzzzzzzzzzzzTDDDzzzzzzzzzzzzzzzzzRRzzzzxwzzzzzzzzzzzzzzzzBzzzzzzhwzjrjzzzzzzzzzzyjzzzzzzQzyo331rzzzzzzzzyzzzzzzyRyD733321rzzzzzzzhTzzzzzhT373/AB8sdO88888880888888s/c5rN9y4823888888888888s8muc6V89999o888884U88888o8888+V999998/88884888888M88888f999999W88888888888888888+999999V88888888888088888z999999V88888888888sc888pe999999V8888888888888IL8c9999999V888888888888sM8OZ8999999V888888888888888889999999V88888888888088888t999999V88888888888U88888+999999V88884888888c888888999999e88888c88888888888+39999t/8APPPPGNPPPBPPPPPPPLNP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",
			},
			Plans: []brokerapi.ServicePlan{
				brokerapi.ServicePlan{
					Name:        "default-apicurio",
					ID:          "default-apicurio",
					Description: "default apicurio plan",
					Free:        true,
					Schemas: &brokerapi.Schemas{
						ServiceBinding: &brokerapi.ServiceBindingSchema{
							Create: &brokerapi.RequestResponseSchema{},
						},
						ServiceInstance: &brokerapi.ServiceInstanceSchema{
							Create: &brokerapi.InputParametersSchema{},
						},
					},
				},
			},
		},
	}
}
