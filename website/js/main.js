const baseUrl = 'http://localhost:1323/api/v1/simulateCombinations'

Vue.use(window.vuelidate.default);
const { required, numeric, maxLength } = window.validators;

const form = new Vue({
    el: '#body',
    data: {
        // #form
        title: '入力フォームバリデーション',
        allParticipants: '36',
        participantsInEachGroup: '6',
        repeatCnt: '3',
        trials: '1000',
        // #table
        simulationResult: {},
    },
    validations: {
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
            maxLength: maxLength(3)
        },
        trials: {
            required,
            maxLength: maxLength(6)
        }
    },
    methods: {
        async submitForm() {
            this.$v.$touch();
            if (this.$v.$invalid) {
                return false;
            }

            const res = await axios.get(baseUrl, {
                params: {
                    allParticipants: this.allParticipants,
                    participantsInEachGroup: this.participantsInEachGroup,
                    repeatCnt: this.repeatCnt,
                    trials: this.trials,
                }
            }).catch(error => {
                return err.response
            });

            console.log(res.data);
            this.simulationResult = res.data;

            if (res.status != 200) {
                console.log(res.status);
                alert("入力値が正しくありません。");
            }
        }
    }
});