const baseUrl = 'http://localhost:1323/api/v1/simulateCombinations'

Vue.use(window.vuelidate.default);
const { required, numeric, maxLength } = window.validators;

const form = new Vue({
    el: 'main',
    data: {
        // #form
        title: '入力フォームバリデーション',
        allParticipants: '36',
        participantsInEachGroup: '6',
        repeatCnt: '3',
        trials: '1000',
        // #result
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
            maxLength: maxLength(2)
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
            }).catch(err => {
                return err.response;
            });

            if(res.status != 200) {
                alert(res.data.message);
                return false;
            }

            console.log(res.data);
            this.simulationResult = res.data;
        }
    }
});