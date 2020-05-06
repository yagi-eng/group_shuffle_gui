const baseUrl = location.href;

Vue.use(window.vuelidate.default);
const { required, numeric, maxLength } = window.validators;

// validation function
const equalsAllParticipants = function (val) {
    if (val == "") {
        return true;
    }

    let numOfComma = (val.match(/,/g) || []).length;
    return main.allParticipants == numOfComma + 1;
}

const main = new Vue({
    el: 'main',
    data: {
        // #form
        title: '入力フォームバリデーション',
        allParticipants: '18',
        participantsInEachGroup: '6',
        repeatCnt: '3',
        trials: '1000',
        names: '田中,山田,佐藤,斎藤,近藤,伊藤,高橋,織田,豊臣,徳永,不破,貴島,深田,滝川,弘中,玉城,小倉,尾崎',
        // #result
        simulationResult: {},
    },
    validations: {
        // #form
        allParticipants: {
            required,
            maxLength: maxLength(3)
        },
        participantsInEachGroup: {
            required,
            maxLength: maxLength(3)
        },
        repeatCnt: {
            required,
            maxLength: maxLength(2)
        },
        trials: {
            required,
            maxLength: maxLength(6)
        },
        names: {
            equalsAllParticipants
        }
    },
    methods: {
        async submitForm() {
            this.$v.$touch();
            if (this.$v.$invalid) {
                return false;
            }

            const res = await axios.get(baseUrl + "api/v1/simulateCombinations", {
                params: {
                    allParticipants: this.allParticipants,
                    participantsInEachGroup: this.participantsInEachGroup,
                    repeatCnt: this.repeatCnt,
                    trials: this.trials,
                    names: this.names,
                }
            }).catch(err => {
                return err.response;
            });

            if (res.status != 200) {
                alert(res.data.message);
                return false;
            }

            console.log(res.data);
            this.simulationResult = res.data;
        }
    }
});