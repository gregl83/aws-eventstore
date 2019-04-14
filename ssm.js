const AWS = require('aws-sdk');
const password = require('generate-password');

AWS.config.update({
    region: 'us-east-1'
});

const ssm = new AWS.SecretsManager();

const passwordConfig = {
    length: 16,
    numbers: true,
    strict: true,
};

const secretConfig = secret => {
    return {
        Description: "EventStore in Aurora database secret created by serverless",
        Name: "EVS_AURORA_PASSWORD",
        SecretString: secret
    };
};

module.exports.AURORA_PASSWORD = () => {
    return ssm.getSecretValue({SecretId: 'EVS_AURORA_PASSWORD'})
        .promise()
        .then(secret => secret.SecretString)
        .catch(err => {
            if (err.name !== 'ResourceNotFoundException') return console.log(err);

            const secret = password.generate(passwordConfig);

            return ssm.createSecret(secretConfig(secret)).promise();
        });
};