/**
 * Builds the expected request body for an ElasticSearch request using the
 * original proxied request body
 *
 * @param {Object} body - The original request body
 * @returns {Object} The request body expected by ElasticSearch
 */
function buildElasticSearchRequestBody(body) {
  if (body.query == null || body.filters == null) {
    return body;
  }

  const { query, filters } = body;

  delete body.query;
  delete body.filters;

  return {
    ...body,
    query: {
      bool: {
        must: {
          query_string: {
            query: `*${sanitizeQuery(query)}*`,
            fields: ['name', 'owner.name', 'description'],
          },
        },
        filter: {
          terms: {
            'vimColorSchemes.backgrounds': filters,
          },
        },
      },
    },
  };
}

/*
 * Sanitizes a query entered by the user in order to safely send it over to
 * the ElasticSearch proxy
 *
 * @param {string} query - The query entered by the used
 * @returns {string} The query with all unsafe characters escaped
 */
function sanitizeQuery(query) {
  if (query == null) {
    return '';
  }

  const characters = query.split('');

  for (let i = 0; i < query.length; i++) {
    characters[i] = sanitizeCharacter(query[i]);
  }

  return characters.join('');
}

function sanitizeCharacter(character) {
  if (character == null) {
    return '';
  }

  return character
    .replace(/[*+\-=~><"?^${}():!/[\]\\\s]/g, '\\$&') // Replace reserved characters
    .replace(/\|\|/g, '\\||') // replace ||
    .replace(/&&/g, '\\&&') // replace &&
    .replace(/AND/g, '\\A\\N\\D') // replace AND
    .replace(/OR/g, '\\O\\R') // replace OR
    .replace(/NOT/g, '\\N\\O\\T'); // replace NOT
}

const ElasticSearchHelper = {
  buildElasticSearchRequestBody,
  sanitizeQuery,
};

module.exports = ElasticSearchHelper;
